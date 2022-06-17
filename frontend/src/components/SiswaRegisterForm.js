import { useNavigate } from 'react-router-dom';
import { registerSiswa } from '../lib/siswa';
import useAuthStore from '../store/auth';
import Input from './Input';

function SiswaRegisterForm({
  formValues,
  handleInputChange,
  showPassword,
  toggleShowPassword,
}) {
  const navigate = useNavigate();
  const setUser = useAuthStore((state) => state.setUser);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const { confirm, ...values } = formValues;
    const data = await registerSiswa(values);
    setUser({ ...data, ...values });
    navigate('/');
  };

  return (
    <form className="space-y-6 bg-gray-200 p-8" onSubmit={handleSubmit}>
      <Input
        name="nama"
        value={formValues.nama || ''}
        type="text"
        label="Nama"
        onChange={handleInputChange}
      />

      <Input
        name="email"
        value={formValues.email || ''}
        type="email"
        label="Alamat email"
        onChange={handleInputChange}
      />

      <div className="grid grid-cols-2 gap-x-6 gap-y-2">
        <Input
          name="password"
          value={formValues.password || ''}
          type={showPassword ? 'text' : 'password'}
          label="Password"
          onChange={handleInputChange}
        />
        <Input
          name="confirm"
          value={formValues.confirm || ''}
          type={showPassword ? 'text' : 'password'}
          label="Konfirmasi password"
          onChange={handleInputChange}
        />
        <div className="ml-1 flex w-fit items-center gap-3">
          <input
            type="checkbox"
            name="showPassword"
            id="showPassword"
            value={showPassword}
            onClick={toggleShowPassword}
            className="block"
          />
          <label htmlFor="showPassword" className="block">
            Lihat password
          </label>
        </div>
      </div>

      <Input
        name="tanggalLahir"
        value={formValues.tanggalLahir || ''}
        type="date"
        label="Tanggal lahir"
        onChange={handleInputChange}
      />

      <Input
        name="alamat"
        value={formValues.alamat || ''}
        type="text"
        label="Alamat"
        onChange={handleInputChange}
      />

      <Input
        name="nomorTelepon"
        value={formValues.nomorTelepon || ''}
        type="text"
        label="Nomor telepon"
        onChange={handleInputChange}
      />

      <Input
        name="namaBank"
        value={formValues.namaBank || ''}
        type="text"
        label="Nama bank"
        onChange={handleInputChange}
      />

      <Input
        name="nomorRekening"
        value={formValues.nomorRekening || ''}
        type="text"
        label="Nomor rekening"
        onChange={handleInputChange}
      />

      <div>
        <label
          htmlFor="tingkatPendidikan"
          className="block text-sm font-medium"
        >
          Tingkat pendidikan
        </label>
        <select
          name="tingkatPendidikan"
          value={formValues.tingkatPendidikan}
          defaultValue="default"
          id="tingkatPendidikan"
          onChange={handleInputChange}
          className="mt-1 block w-full rounded-sm border-gray-400 py-1 px-2"
        >
          <option value="default" disabled>
            Pilih tingkat pendidikan
          </option>
          <option value="SMA">SMA</option>
          <option value="S1">S1</option>
          <option value="S2">S2</option>
        </select>
      </div>

      <Input
        name="namaInstansi"
        value={formValues.namaInstansi || ''}
        type="text"
        label="Sekolah/Universitas"
        onChange={handleInputChange}
      />

      <button className="w-full rounded-sm bg-black py-2 text-white hover:bg-gray-800">
        Submit
      </button>
    </form>
  );
}

export default SiswaRegisterForm;
