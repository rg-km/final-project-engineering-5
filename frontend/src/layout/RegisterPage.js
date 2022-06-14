import { useState } from 'react';
import MitraRegisterForm from '../components/MitraRegisterForm';
import SiswaRegisterForm from '../components/SiswaRegisterForm';

function RegisterPage() {
  const [role, setRole] = useState('siswa');
  const [formValues, setFormValues] = useState({});
  const [showPassword, setShowPassword] = useState(false);

  const handleInputChange = (e) => {
    setFormValues({
      ...formValues,
      [e.target.name]: e.target.value,
    });
  };

  const clearFormValues = () => {
    setFormValues({});
  };

  const toggleShowPassword = () => {
    setShowPassword(!showPassword);
  };

  return (
    <div className="mx-auto max-w-[448px] py-10">
      <div>
        <button
          onClick={() => {
            setRole('siswa');
            clearFormValues();
          }}
          className={`px-4 py-2 ${
            role === 'siswa' ? 'bg-gray-200' : 'bg-white'
          }`}
        >
          Siswa
        </button>
        <button
          onClick={() => {
            setRole('mitra');
            clearFormValues();
          }}
          className={`px-4 py-2 ${
            role === 'mitra' ? 'bg-gray-200' : 'bg-white'
          }`}
        >
          Mitra
        </button>
      </div>

      {role === 'siswa' ? (
        <SiswaRegisterForm
          formValues={formValues}
          handleInputChange={handleInputChange}
          showPassword={showPassword}
          toggleShowPassword={toggleShowPassword}
        />
      ) : role === 'mitra' ? (
        <MitraRegisterForm
          formValues={formValues}
          handleInputChange={handleInputChange}
          showPassword={showPassword}
          toggleShowPassword={toggleShowPassword}
        />
      ) : null}
    </div>
  );
}

export default RegisterPage;
