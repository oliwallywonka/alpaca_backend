import axios from 'axios'

const headers = {
  'Content-Type': 'application/json',
  'Access-Control-Allow-Origin': '*',
}

const instance = axios.create({
  baseURL: 'http://localhost:8000',
  headers,
})

export { instance as API }
