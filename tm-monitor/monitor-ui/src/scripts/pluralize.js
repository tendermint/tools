import pluralize from 'pluralize'

export default function (string, number = 0) {
  return pluralize(string, number)
}
