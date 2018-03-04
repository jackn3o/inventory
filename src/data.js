import moment from 'moment'
import _ from 'lodash'

let categories = [


]

let list = [
  {
    active: false,
    description: 'Toys',
    items: [{
      value: false,
      id: '0001',
      code: 'VIB 01',
      description: 'Vibration',
      color: 'Pink',
      balance: 24
    }]
  },
  {
    active: false,
    description: 'Bras',
    items: [{
      categories: 'Bras',
      value: false,
      id: '0002',
      code: 'Bra No.4',
      description: 'Sporty Bra-Limited Edition',
      color: 'White',
      balance: 24
    }]
  }
]

let itemDetails = [
  {
    itemId: '0001',
    batchNo: 'B001',
    documentNo: '12345',
    date: moment('1/4/2017', 'DD-M-YYYY'),
    in: 100,
    out: null,
    unitCost: 2.0,
    balanceQuantity: 100,
    totalCost: 200.0,
    totalSales: null,
    balanceCost: 200.0
  },
  {
    itemId: '0001',
    batchNo: 'B003',
    documentNo: '12346',
    date: moment('3/5/2017', 'DD-M-YYYY'),
    in: 50,
    out: null,
    unitCost: 2.1,
    balanceQuantity: 130,
    totalCost: 105.0,
    totalSales: null,
    balanceCost: 265.0
  },
  {
    itemId: '0001',
    batchNo: 'B005',
    documentNo: '12350',
    date: moment('10/5/2017', 'DD-M-YYYY'),
    in: 20,
    out: null,
    unitCost: 1.85,
    balanceQuantity: 150,
    totalCost: 37.0,
    totalSales: null,
    balanceCost: 302.0
  },
  {
    itemId: '0001',
    batchNo: null,
    documentNo: 'INV20170430',
    date: moment('30/4/2017', 'DD-M-YYYY'),
    in: null,
    out: 20,
    unitCost: 4.9,
    balanceQuantity: 80,
    totalCost: null,
    totalSales: 98.0,
    balanceCost: 160.0
  },
  {
    itemId: '0001',
    batchNo: null,
    documentNo: 'INV20170531',
    date: moment('31/5/2017', 'DD-M-YYYY'),
    in: null,
    out: 50,
    unitCost: 4.9,
    balanceQuantity: 100,
    totalCost: null,
    totalSales: 245.0,
    balanceCost: 202.0
  },
  {
    itemId: '0001',
    batchNo: null,
    documentNo: 'INV20170630',
    date: moment('30/6/2017', 'DD-M-YYYY'),
    in: null,
    out: 10,
    unitCost: 4.9,
    balanceQuantity: 90,
    totalCost: null,
    totalSales: 49.0,
    balanceCost: 182.0
  },
  {
    itemId: '0001',
    batchNo: null,
    documentNo: 'INV20170731',
    date: moment('31/7/2017', 'DD-M-YYYY'),
    in: null,
    out: 25,
    unitCost: 4.9,
    balanceQuantity: 65,
    totalCost: null,
    totalSales: 122.5,
    balanceCost: 131.5
  },
  {
    itemId: '0001',
    batchNo: null,
    documentNo: 'INV20170831',
    date: moment('31/8/2017', 'DD-M-YYYY'),
    in: null,
    out: 35,
    unitCost: 4.9,
    balanceQuantity: 30,
    totalCost: null,
    totalSales: 171.5,
    balanceCost: 68.5
  },
  {
    itemId: '0002',
    batchNo: 'B002',
    documentNo: '12345',
    date: moment('1/4/2017', 'DD-M-YYYY'),
    in: 200,
    out: null,
    unitCost: 2.55,
    balanceQuantity: 100,
    totalCost: 510.0,
    totalSales: null,
    balanceCost: 200.0
  },
  {
    itemId: '0002',
    batchNo: null,
    documentNo: 'INV20170430',
    date: moment('30/4/2017', 'DD-M-YYYY'),
    in: null,
    out: 20,
    unitCost: 5.1,
    balanceQuantity: 80,
    totalCost: null,
    totalSales: 102.0,
    balanceCost: 204.0
  },
  {
    itemId: '0002',
    batchNo: 'B004',
    documentNo: '12346',
    date: moment('3/5/2017', 'DD-M-YYYY'),
    in: 100,
    out: null,
    unitCost: 2.65,
    balanceQuantity: 180,
    totalCost: 265.0,
    totalSales: null,
    balanceCost: 469.0
  },
  {
    itemId: '0002',
    batchNo: 'B006',
    documentNo: '12350',
    date: moment('10/5/2017', 'DD-M-YYYY'),
    in: 100,
    out: null,
    unitCost: 2.4,
    balanceQuantity: 280,
    totalCost: 240.0,
    totalSales: null,
    balanceCost: 709.0
  },
  {
    itemId: '0002',
    batchNo: null,
    documentNo: 'INV20170531',
    date: moment('31/5/2017', 'DD-M-YYYY'),
    in: null,
    out: 50,
    unitCost: 5.1,
    balanceQuantity: 230,
    totalCost: null,
    totalSales: 255.0,
    balanceCost: 581.5
  },
  {
    itemId: '0002',
    batchNo: null,
    documentNo: 'INV20170630',
    date: moment('30/6/2017', 'DD-M-YYYY'),
    in: null,
    out: 20,
    unitCost: 5.1,
    balanceQuantity: 210,
    totalCost: null,
    totalSales: 102.0,
    balanceCost: 530.5
  },
  {
    itemId: '0002',
    batchNo: null,
    documentNo: 'INV20170731',
    date: moment('31/7/2017', 'DD-M-YYYY'),
    in: null,
    out: 5,
    unitCost: 5.1,
    balanceQuantity: 205,
    totalCost: null,
    totalSales: 25.5,
    balanceCost: 517.75
  },
  {
    itemId: '0002',
    batchNo: null,
    documentNo: 'INV20170831',
    date: moment('31/8/2017', 'DD-M-YYYY'),
    in: null,
    out: 30,
    unitCost: 4.9,
    balanceQuantity: 175,
    totalCost: null,
    totalSales: 147.0,
    balanceCost: 438.75
  }
]

export let getInventory = () =>
  new Promise(resolve => {
    list = _.sortBy(list, obj_item => {
      return obj_item.description
    })
    resolve(list)
  })

export let getInventoryById = str_id =>
  new Promise(resolve => {
    resolve(
      list.find(obj_item => {
        return obj_item.id == str_id
      })
    )
  })

export let addNewItem = obj_item =>
  new Promise(resolve => {
    list.push(obj_item)
    resolve(list)
  })

export let getInventoryDetailById = str_id =>
  new Promise(resolve => {
    let currentDetail = itemDetails.filter(obj_detail => {
      return obj_detail.itemId == str_id
    })

    currentDetail = _.sortBy(currentDetail, obj_detail => {
      return obj_detail.date
    }).reverse()

    resolve(currentDetail)
  })
