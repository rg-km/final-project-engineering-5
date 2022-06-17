import axios from 'axios';

export async function login(email, password) {
  try {
    const response = await axios.post(
      `${process.env.REACT_APP_API_BASE_URL}/api/login`,
      { email, password }
    );
    return response.data;
  } catch (error) {
    console.log(error);
  }
}
