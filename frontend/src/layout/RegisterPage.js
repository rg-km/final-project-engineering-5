import { useState } from 'react';
import MitraRegisterForm from '../components/MitraRegisterForm';
import SiswaRegisterForm from '../components/SiswaRegisterForm';

function RegisterPage() {
  const [role, setRole] = useState('siswa');

  return (
    <div className="mx-auto max-w-[448px] py-10">
      <div>
        <button
          onClick={() => {
            setRole('siswa');
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
          }}
          className={`px-4 py-2 ${
            role === 'mitra' ? 'bg-gray-200' : 'bg-white'
          }`}
        >
          Mitra
        </button>
      </div>

      {role === 'siswa' ? (
        <SiswaRegisterForm />
      ) : role === 'mitra' ? (
        <MitraRegisterForm />
      ) : null}
    </div>
  );
}

export default RegisterPage;
