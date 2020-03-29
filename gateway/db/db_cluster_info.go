package db

import "time"

type ClusterInfo struct {
	ID          int64  `gorm:"column:id;primary_key"`
	Name        string `gorm:"type:varchar(64);unique;not null;column:name"`
	Description string `gorm:"type:varchar(4096);column:description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (obj *ClusterInfo) TableName() string {
	return TablePrefix + "cluster_info"
}

func (obj *ClusterInfo) New() error {
	return db.Save(obj).Error
}

func (obj *ClusterInfo) Modify() error {
	return db.Table(obj.TableName()).Where(QueryID, &obj.ID).Update(obj).Error
}

func (obj *ClusterInfo) QueryByID(_id *int64) error {
	return db.Table(obj.TableName()).Where(QueryID, *_id).First(obj).Error
}

func (obj *ClusterInfo) Del() error {
	if err := db.Table(obj.TableName()).Where(QueryID, &obj.ID).First(obj).Error; err != nil {
		return err
	}
	return db.Where(QueryID, &obj.ID).Delete(ClusterInfo{}).Error
}
