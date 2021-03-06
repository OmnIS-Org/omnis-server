package db

import (
	"database/sql"
	"fmt"

	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

// GetInterfaces should have a comment.
func GetInterfaces(automatic bool) (model.InterfaceOs, error) {
	log.Debug(fmt.Sprintf("GetInterfaceOs(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_interfaces(?);", automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var interfaceOs model.InterfaceOs

	for rows.Next() {
		var interfaceO model.InterfaceO

		err := rows.Scan(&interfaceO.ID, &interfaceO.Name, &interfaceO.Ipv4, &interfaceO.Ipv4Mask, &interfaceO.MAC, &interfaceO.InterfaceType, &interfaceO.MachineID, &interfaceO.NetworkID)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		interfaceOs = append(interfaceOs, interfaceO)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return interfaceOs, nil
}

// GetInterface should have a comment.
func GetInterface(id int32, automatic bool) (*model.InterfaceO, error) {
	log.Debug(fmt.Sprintf("GetInterfaceO(%d,%t)", id, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var interfaceO model.InterfaceO
	err = db.QueryRow("CALL get_interface_by_id(?,?);", id, automatic).Scan(&interfaceO.ID, &interfaceO.Name, &interfaceO.Ipv4, &interfaceO.Ipv4Mask, &interfaceO.MAC, &interfaceO.InterfaceType, &interfaceO.MachineID, &interfaceO.NetworkID)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &interfaceO, nil
}

// InsertInterface should have a comment.
func InsertInterface(interfaceO *model.InterfaceO, automatic bool) (int32, error) {
	log.Debug(fmt.Sprintf("InsertInterfaceO(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var id int32 = 0
	sqlStr := "CALL insert_interface(?,?,?,?,?,?,?,?);"

	err = db.QueryRow(sqlStr, interfaceO.Name, interfaceO.Ipv4, interfaceO.Ipv4Mask, interfaceO.MAC, interfaceO.InterfaceType, interfaceO.MachineID, interfaceO.NetworkID, automatic).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return id, nil
}

// UpdateInterface should have a comment.
func UpdateInterface(id int32, interfaceO *model.InterfaceO, automatic bool) (int64, error) {
	log.Debug(fmt.Sprintf("UpdateInterfaceO(%t)", automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	sqlStr := "CALL update_interface(?,?,?,?,?,?,?,?,?);"

	res, err := db.Exec(sqlStr, id, interfaceO.Name, interfaceO.Ipv4, interfaceO.Ipv4Mask, interfaceO.MAC, interfaceO.InterfaceType, interfaceO.MachineID, interfaceO.NetworkID, automatic)

	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// DeleteInterface should have a comment.
func DeleteInterface(id int32) (int64, error) {
	log.Debug(fmt.Sprintf("DeleteInterfaceO(%d)", id))

	db, err := GetOmnisConnection()
	if err != nil {
		return 0, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	res, err := db.Exec("CALL delete_interface(?);", id)
	if err != nil {
		return 0, fmt.Errorf("db.Exec failed <- %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowsAffected failed <- %v", err)
	}

	return rowsAffected, nil
}

// GetInterfaceByMac should have a comment.
func GetInterfaceByMac(mac string, automatic bool) (*model.InterfaceO, error) {
	log.Debug(fmt.Sprintf("GetInterfaceOByMac(%s,%t)", mac, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	var interfaceO model.InterfaceO
	err = db.QueryRow("CALL get_interface_by_mac(?,?);", mac, automatic).Scan(&interfaceO.ID, &interfaceO.Name, &interfaceO.Ipv4, &interfaceO.Ipv4Mask, &interfaceO.MAC, &interfaceO.InterfaceType, &interfaceO.MachineID, &interfaceO.NetworkID)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("db.QueryRow failed <- %v", err)
	}

	return &interfaceO, nil
}

// GetInterfacesByMachineID should have a comment.
func GetInterfacesByMachineID(machineID int32, automatic bool) (model.InterfaceOs, error) {
	log.Debug(fmt.Sprintf("GetInterfaceOsByMachineId(%d,%t)", machineID, automatic))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_interfaces_by_machine_id(?,?);", machineID, automatic)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var interfaceOs model.InterfaceOs

	for rows.Next() {
		var interfaceO model.InterfaceO

		err := rows.Scan(&interfaceO.ID, &interfaceO.Name, &interfaceO.Ipv4, &interfaceO.Ipv4Mask, &interfaceO.MAC, &interfaceO.InterfaceType, &interfaceO.MachineID, &interfaceO.NetworkID)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		interfaceOs = append(interfaceOs, interfaceO)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return interfaceOs, nil
}

// GetOutdatedInterfaces only return authorized machines
func GetOutdatedInterfaces(outdatedDay int) (model.InterfaceOs, error) {
	log.Debug(fmt.Sprintf("GetMachines(%d)", outdatedDay))

	db, err := GetOmnisConnection()
	if err != nil {
		return nil, fmt.Errorf("GetOmnisConnection failed <- %v", err)
	}

	rows, err := db.Query("CALL get_outdated_interfaces(?);", outdatedDay)
	if err != nil {
		return nil, fmt.Errorf("db.Query failed <- %v", err)
	}
	defer rows.Close()

	var interfaceOs model.InterfaceOs

	for rows.Next() {
		var interfaceO model.InterfaceO

		err := rows.Scan(&interfaceO.ID,
			&interfaceO.Name,
			&interfaceO.Ipv4,
			&interfaceO.Ipv4Mask,
			&interfaceO.MAC,
			&interfaceO.InterfaceType,
			&interfaceO.MachineID,
			&interfaceO.NetworkID,
			&interfaceO.NameLastModification,
			&interfaceO.Ipv4LastModification,
			&interfaceO.Ipv4MaskLastModification,
			&interfaceO.MACLastModification,
			&interfaceO.InterfaceTypeLastModification,
			&interfaceO.MachineIDLastModification,
			&interfaceO.NetworkIDLastModification)

		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed <- %v", err)
		}

		interfaceOs = append(interfaceOs, interfaceO)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Scan failed <- %v", err)
	}

	return interfaceOs, nil
}

// GetInterfacesO should have a comment.
func GetInterfacesO(automatic bool) (model.Objects, error) {
	return GetInterfaces(automatic)
}

// GetInterfaceO should have a comment.
func GetInterfaceO(id int32, automatic bool) (model.Object, error) {
	return GetInterface(id, automatic)
}

// InsertInterfaceO should have a comment.
func InsertInterfaceO(object *model.Object, automatic bool) (int32, error) {
	var interfaceO *model.InterfaceO = (*object).(*model.InterfaceO)
	return InsertInterface(interfaceO, automatic)
}

// UpdateInterfaceO should have a comment.
func UpdateInterfaceO(id int32, object *model.Object, automatic bool) (int64, error) {
	var interfaceO *model.InterfaceO = (*object).(*model.InterfaceO)
	return UpdateInterface(id, interfaceO, automatic)
}

// GetInterfaceByMacO should have a comment.
func GetInterfaceByMacO(mac string, automatic bool) (model.Object, error) {
	return GetInterfaceByMac(mac, automatic)
}

// GetInterfacesByMachineIDO should have a comment.
func GetInterfacesByMachineIDO(machineID int32, automatic bool) (model.Objects, error) {
	return GetInterfacesByMachineID(machineID, automatic)
}

// GetOutdatedInterfacesO should have a comment.
func GetOutdatedInterfacesO(outdatedDay int) (model.Objects, error) {
	return GetOutdatedInterfaces(outdatedDay)
}
