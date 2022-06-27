import axios from 'axios';

export async function login(role, email, password) {
  const response = await axios.post(
    `${process.env.REACT_APP_API_BASE_URL}/api/${role.toLowerCase()}/login`,
    {},
    { auth: { username: email, password } }
  );
  return response.data;
}
