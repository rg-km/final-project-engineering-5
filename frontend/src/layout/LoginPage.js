import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Input from '../components/Input';
import useAuthStore from '../store/auth';
import { login } from '../lib/login';

function LoginPage() {
  const [formValues, setFormValues] = useState({});
  const [role, setRole] = useState('SISWA');
  const navigate = useNavigate();
  const setUser = useAuthStore((state) => state.setUser);

  const handleInputChange = (e) => {
    setFormValues({
      ...formValues,
      [e.target.name]: e.target.value,
    });
  };

  const clearFormValues = () => {
    setFormValues({});
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const data = await login(role, formValues.email, formValues.password);
    setUser(data);
    navigate('/dashboard');
  };

  return (
    <div className="mx-auto max-w-[448px] py-10">
      <div>
        <button
          onClick={() => {
            setRole('SISWA');
            clearFormValues();
          }}
          className={`px-4 py-2 ${
            role === 'SISWA' ? 'bg-gray-200' : 'bg-white'
          }`}
        >
          Siswa
        </button>
        <button
          onClick={() => {
            setRole('MITRA');
            clearFormValues();
          }}
          className={`px-4 py-2 ${
            role === 'MITRA' ? 'bg-gray-200' : 'bg-white'
          }`}
        >
          Mitra
        </button>
      </div>
      <form className="space-y-6 bg-gray-200 p-8" onSubmit={handleSubmit}>
        <Input
          name="email"
          value={formValues.email || ''}
          type="email"
          label="Alamat email"
          onChange={handleInputChange}
        />
        <Input
          name="password"
          value={formValues.password || ''}
          type="password"
          label="Password"
          onChange={handleInputChange}
        />

        <button className="w-full rounded-sm bg-black py-2 text-white hover:bg-gray-800">
          Submit
        </button>
      </form>
    </div>
  );
}

export default LoginPage;
