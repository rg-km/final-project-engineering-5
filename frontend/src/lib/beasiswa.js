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
