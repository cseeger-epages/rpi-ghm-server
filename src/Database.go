/*
   GOLANG REST API Skeleton

   Copyright (C) 2017 Carsten Seeger

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.

   @author Carsten Seeger
   @copyright Copyright (C) 2017 Carsten Seeger
   @license http://www.gnu.org/licenses/gpl-3.0 GNU General Public License 3
   @link https://github.com/cseeger-epages/rest-api-go-skeleton
*/

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func DBOpen() (*sql.DB, error) {
	system := fmt.Sprintf("%s:%s", Conf.DB.Host, Conf.DB.Port)
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s", Conf.DB.User, Conf.DB.Password, "tcp", system, Conf.DB.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func mysqldummy() (Projects, error) {
	var (
		id  int
		val string
		ret []Project
	)

	db, err := DBOpen()
	defer db.Close()

	if err != nil {
		return Projects{}, err
	}

	data, err := db.Query("select project_id, value from t_project")
	defer data.Close()
	if err != nil {
		return Projects{}, err
	}

	for data.Next() {
		err = data.Scan(&id, &val)
		if err != nil {
			return Projects{}, err
		}
		ret = append(ret, Project{Id: id, Name: val})
	}

	return Projects{ret}, nil
}
