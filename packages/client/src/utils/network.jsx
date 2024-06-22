import axios from 'axios'

const getAssets = async (page = 1, pageSize = 10, type = null) => {
  console.log('aaaaa', type)
  try {
    const token = localStorage.getItem('token')

    if (!token) {
      throw Error('missing token')
    }
    const res = await axios.get(
      `http://localhost:3000/api/v1/assets${
        type ? `/${type}` : ''
      }?page=${page}&limit=${pageSize}`,
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      }
    )
    console.log('aaa', res.data)
    return res.data
  } catch (e) {
    throw new Error(e.message)
  }
}

const getUserFavorites = async (userId, page = 1, pageSize = 10) => {
  try {
    const token = localStorage.getItem('token')

    if (!token) {
      throw Error('missing token')
    }
    const res = await axios.get(
      `http://localhost:3000/api/v1/users/${userId}/favorites?page=${page}&limit=${pageSize}`,
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      }
    )
    console.log('aaa', res.data)
    return res.data
  } catch (e) {
    throw new Error(e.message)
  }
}

export { getAssets, getUserFavorites }
