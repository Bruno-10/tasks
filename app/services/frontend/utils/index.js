import { subWeeks, addDays, subDays, format, formatISO } from 'date-fns'
const today = () => formatISO(new Date(), { representation: 'date' })
const tomorrow = () =>
  formatISO(addDays(new Date(), 1), { representation: 'date' })
const weekAgo = () =>
  formatISO(subWeeks(new Date(), 1), { representation: 'date' })
const monthAgo = () =>
  formatISO(subDays(new Date(), 30), { representation: 'date' })
const formatMonth = (date = new Date()) => {
  const dateInstance = new Date(date)

  return format(dateInstance, 'yyyy-MM')
}

const formattedDate = (str) => {
  if (!str) return ''

  if (typeof str === 'string') {
    return format(new Date(str), 'dd-MM-yyyy')
  }

  return ''
}

const formattedDateWithTime = (str) => {
  if (!str) return ''

  if (typeof str === 'string') {
    return format(new Date(str), 'dd-MM-yyyy HH:mm:ss')
  }

  return ''
}

const capitalizeFirstLetter = (string) => {
  if (!string) return ''

  if (typeof string === 'string') {
    return string.charAt(0).toUpperCase() + string.slice(1)
  }

  return ''
}

export default {
  today,
  tomorrow,
  weekAgo,
  monthAgo,
  formattedDate,
  formattedDateWithTime,
  formatMonth,
  capitalizeFirstLetter,
}
