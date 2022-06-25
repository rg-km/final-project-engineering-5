import axios from 'axios';

export async function registerSiswa(values) {
  try {
    const response = await axios.post(
      `${process.env.REACT_APP_API_BASE_URL}/api/siswa/signup`,
      values
    );
    return response.data;
  } catch (error) {
    console.log(error);
  }
}

export async function getSiswa(token) {
  try {
    const response = await axios.get(
      `${process.env.REACT_APP_API_BASE_URL}/api/siswa/detail`,
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response.data;
  } catch (error) {
    console.log(error);
  }
}
