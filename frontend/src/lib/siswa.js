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
