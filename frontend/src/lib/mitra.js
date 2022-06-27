import axios from 'axios';

export async function registerMitra(values) {
  try {
    const response = await axios.post(
      `${process.env.REACT_APP_API_BASE_URL}/api/mitra/signup`,
      values
    );
    return response.data;
  } catch (error) {
    console.log(error);
  }
}

export async function getListBeasiswaSiswa(token) {
  const response = await axios.get(
    `${process.env.REACT_APP_API_BASE_URL}/api/beasiswa-siswa`,
    { headers: { Authorization: `Bearer ${token}` } }
  );
  return response.data;
}
