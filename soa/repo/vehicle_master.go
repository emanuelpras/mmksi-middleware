package repo

import (
	"database/sql"
	"log"
	"middleware-mmksi/soa/response"
)

type Pagination struct {
	Page    int
	Limit   int
	Counter int
}

type SoaRepo interface {
	VehicleMasterList(request Pagination) (*response.VehicleMasterData, error)
}

type soaRepo struct {
	DB *sql.DB
}

func NewSoaRepo(db *sql.DB) SoaRepo {
	return &soaRepo{
		DB: db,
	}
}

func (r *soaRepo) VehicleMasterList(request Pagination) (*response.VehicleMasterData, error) {
	// sqlQuery, err := r.DB.Query("select * from vehicle_master limit ? ?", request.Counter, request.Limit)
	// sqlQuery, err := r.DB.Query("select brand from vehicle_master")
	sqlStatement := "select ID from vehicle_master"
	row := r.DB.QueryRow(sqlStatement)

	var datas response.VehicleMasterData
	// err := row.Scan(&datas.ID, &datas.Brand, &datas.Model, &datas.VehicleName, &datas.DsfAssetCode, &datas.MmksiType, &datas.MmksiColor, &datas.Package, &datas.DpMinMax)
	err := row.Scan(datas.ID)
	test := r.DB.Ping()
	log.Println(" >>>>>>> pingnya >>>> ", test)
	if err != nil {
		return nil, err
	}

	return nil, nil
	/* var datas []datasponse.VehicleMasterData
	return datas, nil */
}
