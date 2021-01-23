package ping

import (
	mssql "SendGrid-with-Node.js/database/mssql"
)

func logHeartbeat(hb heartbeatModel) (err error) {

	if err = mssql.DB.Table("log_heartbeat").Save(hb).Error; err != nil {
		return
	}

	return
}
