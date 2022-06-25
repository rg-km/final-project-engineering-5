import { useNavigate } from 'react-router-dom';
import { registerMitra } from '../lib/mitra';
import { checkPassword } from '../lib/utils';
import useAuthStore from '../store/auth';
import Input from './Input';

function MitraRegisterForm({
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

    if (!checkPassword(values.password, confirm)) return;
    const data = await registerMitra(values);
    setUser(data);
    navigate('/');
  };

  return (
    <form className="space-y-6 bg-gray-200 p-8" onSubmit={handleSubmit}>
      <Input
        name="nama"
        value={formValues.nama || ''}
        type="text"
        label="Nama mitra"
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
      <div>
        <label htmlFor="about" className="block text-sm font-medium">
          About
        </label>
        <textarea
          name="about"
          value={formValues.about || ''}
          id="about"
          rows={5}
          className="mt-1 block w-full rounded-sm border-gray-400 py-1 px-2"
          onChange={handleInputChange}
        ></textarea>
      </div>
      <Input
        name="situs"
        value={formValues.situs || ''}
        type="text"
        label="Situs"
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
        name="namaPic"
        value={formValues.namaPic || ''}
        type="text"
        label="Person in Charge (PIC)"
        onChange={handleInputChange}
      />
      <Input
        name="nomorPic"
        value={formValues.nomorPic || ''}
        type="text"
        label="Nomor PIC"
        onChange={handleInputChange}
      />
      <button className="w-full rounded-sm bg-black py-2 text-white hover:bg-gray-800">
        Submit
      </button>
    </form>
  );
}

export default MitraRegisterForm;
