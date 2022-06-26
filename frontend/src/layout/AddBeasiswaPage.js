import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Input from '../components/Input';
import { addBeasiswa } from '../lib/beasiswa';
import useAuthStore from '../store/auth';

function AddBeasiswaPage() {
  const navigate = useNavigate();
  const user = useAuthStore((state) => state.user);
  const [formValues, setFormValues] = useState({});
  const [error, setError] = useState('');

  const handleInputChange = (e) => {
    setFormValues({
      ...formValues,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await addBeasiswa(user.token, formValues);
      navigate('/dashboard');
    } catch (error) {
      setError(error.message);
    }
  };

  return (
    <div className="mx-auto max-w-[448px] py-10">
      <form className="space-y-6 bg-gray-200 p-8" onSubmit={handleSubmit}>
        <Input
          name="judulBeasiswa"
          value={formValues.judulBeasiswa || ''}
          type="text"
          label="Judul Beasiswa"
          onChange={handleInputChange}
        />

        <div>
          <label htmlFor="about" className="block text-sm font-medium">
            Deskripsi
          </label>
          <textarea
            name="deskripsi"
            value={formValues.deskripsi || ''}
            id="deskripsi"
            rows={5}
            className="mt-1 block w-full rounded-sm border-gray-400 py-1 px-2"
            onChange={handleInputChange}
          ></textarea>
        </div>

        <Input
          name="tanggalPembukaan"
          value={formValues.tanggalPembukaan || ''}
          type="date"
          label="Tanggal pembukaan"
          onChange={handleInputChange}
        />

        <Input
          name="tanggalPenutupan"
          value={formValues.tanggalPenutupan || ''}
          type="date"
          label="Tanggal penutupan"
          onChange={handleInputChange}
        />

        <div>
          <label htmlFor="about" className="block text-sm font-medium">
            Benefits
          </label>
          <textarea
            name="benefits"
            value={formValues.benefits || ''}
            id="benefits"
            rows={5}
            className="mt-1 block w-full rounded-sm border-gray-400 py-1 px-2"
            onChange={handleInputChange}
          ></textarea>
        </div>

        {error && <p className="font-medium text-red-500">{error}</p>}
        <button className="w-full rounded-sm bg-black py-2 text-white hover:bg-gray-800">
          Submit
        </button>
      </form>
    </div>
  );
}

export default AddBeasiswaPage;
