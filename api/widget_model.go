package main

import (
    "database/sql"
    //"fmt"
    //"errors"
)

type widget struct {
    ID    	   int       `json:"id"`
    Name  	   string    `json:"name"`
    Color      string    `json:"color"`
    Price      float64   `json:"price"`
    Inventory  int       `json:"inventory"`
    Melts      bool      `json:"melts"`
}

func (w *widget) getWidget(db *sql.DB) error {
  return db.QueryRow("SELECT name, color, price, inventory, melts FROM widgets WHERE id=$1", 
  	w.ID).Scan(&w.Name, &w.Color, &w.Price, &w.Inventory, &w.Melts)
}

func (w *widget) updateWidget(db *sql.DB) error {
    _, err :=
        db.Exec("UPDATE widgets SET name=$1, color=$2, price=$3, inventory=$4, melts=$5 WHERE id=$6",
            w.Name, w.Color, w.Price, w.Inventory, w.Melts, w.ID)

    return err
}

func (w *widget) createWidget(db *sql.DB) error {

    res, err := db.Exec("INSERT INTO widgets(name, color, price, inventory, melts) VALUES($1, $2, $3, $4, $5)",
        w.Name, w.Color, w.Price, w.Inventory, w.Melts)

    if err != nil {
        return err
    } 

    id, err := res.LastInsertId()

    if err != nil {
      return err
    } 
      
    w.ID = int(id)
    
    return nil   
}

func getWidgets(db *sql.DB) ([]widget, error) {

  rows, err := db.Query("SELECT * FROM widgets")

  if err != nil {
  	return nil, err
  }

  defer rows.Close()

  widgets := []widget{}

  for rows.Next() {
  	var w widget
  	if err := rows.Scan(&w.ID, &w.Name, &w.Color, &w.Price, &w.Inventory, &w.Melts); err != nil {
  		return nil, err
  	}
  	widgets = append(widgets, w)
  }

  return widgets, nil
}