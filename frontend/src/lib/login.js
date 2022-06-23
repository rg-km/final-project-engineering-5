import axios from 'axios';

// export async function login(email, password) {
//   try {
//     const response = await axios.post(
//       `${process.env.REACT_APP_API_BASE_URL}/api/login`,
//       { email, password }
//     );
//     return response.data;
//   } catch (error) {
//     console.log(error);
//   }
// }

let base64 = require('base-64');

const mitraLogin = async (email, password) => {
  let data = '';

  let config = {
    method: 'post',
    url: `${process.env.REACT_APP_API_BASE_URL}/api/mitra/login`,
    headers: {
      Authorization: `Basic ${base64.encode(email + ':' + password)}`,
    },
    data: data,
  };

  axios(config)
    .then(function (response) {
      console.log(JSON.stringify(response.data));
    })
    .catch(function (error) {
      console.log(error);
    });
};

const siswaLogin = async (email, password) => {
  let data = '';

  let config = {
    method: 'post',
    url: `${process.env.REACT_APP_API_BASE_URL}/api/siswa/login`,
    headers: {
      Authorization: `Basic ${base64.encode(email + ':' + password)}`,
    },
    data: data,
  };

  axios(config)
    .then(function (response) {
      console.log(JSON.stringify(response.data));
    })
    .catch(function (error) {
      console.log(error);
    });
};

export { siswaLogin, mitraLogin };
