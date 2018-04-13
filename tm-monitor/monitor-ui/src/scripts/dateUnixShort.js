import moment from 'moment'
export default function (date) {
  return moment(date, 'x').format('YYYY-MM-DD HH:mm')
}
