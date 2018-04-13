export default function (type, id) {
  var netmonUrl = 'http://91.121.116.55:46670/'
  switch (type) {
    case 'status':
      return netmonUrl + 'status'
    case 'bc':
      return netmonUrl + 'get_chain?chainID=%22' + id + '%22'
    case 'vs':
      return netmonUrl + 'validator_set?valsetID=%22' + id + '%22'
    default:
      return netmonUrl + 'status'
  }
}
