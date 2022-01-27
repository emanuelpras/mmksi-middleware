package repo

import (
	"database/sql"
	"middleware-mmksi/soa/response"
)

type Pagination struct {
	Page    int
	Limit   int
	Counter int
}

type SoaRepo interface {
	VehicleMasterList(request Pagination) (*[]response.VehicleMasterData, int, error)
}

type soaRepo struct {
	DB *sql.DB
}

func NewSoaRepo(db *sql.DB) SoaRepo {
	return &soaRepo{
		DB: db,
	}
}

func (r *soaRepo) VehicleMasterList(request Pagination) (*[]response.VehicleMasterData, int, error) {
	queryData, err := r.DB.Query("select id, brand, model, vehicle_name, dsf_asset_code, mmksi_type, mmksi_color, package, dp_min_max from vehicle_master LIMIT ?, ?", request.Counter, request.Limit)
	if err != nil {
		return nil, 0, err
	}

	sqlStatement := "select count(*) as counter from vehicle_master"
	row := r.DB.QueryRow(sqlStatement)
	var rowCount int
	errCount := row.Scan(&rowCount)

	if errCount != nil {
		return nil, 0, errCount
	}

	var datas []response.VehicleMasterData

	for queryData.Next() {
		var res response.VehicleMasterData
		_ = queryData.Scan(&res.ID, &res.Brand, &res.Model, &res.VehicleName, &res.DsfAssetCode, &res.MmksiType, &res.MmksiColor, &res.Package, &res.DpMinMax)
		datas = append(datas, res)
	}

	return &datas, rowCount, nil
}
