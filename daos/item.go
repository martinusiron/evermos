package daos

import (
	"strconv"

	"github.com/martinusiron/evermos/app"
	"github.com/martinusiron/evermos/models"
)

type ItemDAO struct{}

func NewItemDAO() *ItemDAO {
	return &ItemDAO{}
}

func (dao *ItemDAO) Get(rs app.RequestScope, id int) (*models.Item, error) {

	cachedItem, err := models.FindItem(strconv.Itoa(id))
	if err == models.ErrNoItem {
		var item models.Item
		err := rs.Tx().Select().Model(id, &item)
		models.CacheItem(&item)
		return &item, err
	} else if err != nil {
		return nil, err
	} else {
		return cachedItem, err
	}
}

func (dao *ItemDAO) Create(rs app.RequestScope, item *models.Item) error {
	return rs.Tx().Model(item).Insert()
}

func (dao *ItemDAO) Update(rs app.RequestScope, id int, item *models.Item) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	item.Item_id = id
	return rs.Tx().Model(item).Exclude("Id").Update()
}

func (dao *ItemDAO) Delete(rs app.RequestScope, id int) error {
	item, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(item).Delete()
}

func (dao *ItemDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("item").Row(&count)
	return count, err
}

func (dao *ItemDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Item, error) {
	item := []models.Item{}
	err := rs.Tx().Select().OrderBy("item_id").Offset(int64(offset)).Limit(int64(limit)).All(&item)
	return item, err
}
