import axios from 'axios';

export async function getBeasiswaList(page, limit) {
  try {
    const params = new URLSearchParams({
      page: page ?? 1,
      limit: limit ?? 10,
    });
    const response = await axios.get(
      `${process.env.REACT_APP_API_BASE_URL}/api/beasiswa?${params.toString()}`
    );
    return response.data;
  } catch (error) {
    console.log(error);
  }
}

export async function getBeasiswa(idBeasiswa) {
  const response = await axios.get(
    `${process.env.REACT_APP_API_BASE_URL}/api/beasiswa/${idBeasiswa}`
  );
  return response.data;
}

export async function addBeasiswa(token, values) {
  const response = await axios.post(
    `${process.env.REACT_APP_API_BASE_URL}/api/beasiswa`,
    { ...values },
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );
  return response.data;
}

export async function applyBeasiswa(token, idSiswa, idBeasiswa) {
  const response = await axios.post(
    `${process.env.REACT_APP_API_BASE_URL}/api/beasiswa-siswa`,
    { idSiswa, idBeasiswa },
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );
  return response.data;
}
