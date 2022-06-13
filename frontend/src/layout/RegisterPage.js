import Input from '../components/Input';

function RegisterPage() {
  return (
    <div className="mx-auto max-w-[448px] py-10">
      <div>
        <button className="bg-gray-200 px-4 py-2">Siswa</button>
        <button className="bg-white px-4 py-2">Mitra</button>
      </div>

      <form className="space-y-6 bg-gray-200 p-8">
        <div className="grid grid-cols-2 gap-6">
          <Input name="firstName" type="text" label="Nama depan" />
          <Input name="lastName" type="text" label="Nama belakang" />
        </div>

        <Input name="email" type="email" label="Alamat email" />

        <div className="grid grid-cols-2 gap-x-6 gap-y-2">
          <Input name="password" type="password" label="Password" />
          <Input name="confirm" type="password" label="Konfirmasi password" />
          <div className="ml-1 flex w-fit items-center gap-3">
            <input
              type="checkbox"
              name="showPassword"
              id="showPassword"
              className="block"
            />
            <label htmlFor="showPassword" className="block">
              Lihat password
            </label>
          </div>
        </div>

        <Input name="birthday" type="date" label="Tanggal lahir" />

        <Input name="address" type="text" label="Alamat" />

        <Input name="phoneNumber" type="text" label="Nomor telepon" />

        <Input name="accountNumber" type="text" label="Nomor rekening" />

        <div>
          <label htmlFor="education" className="block text-sm font-medium">
            Tingkat pendidikan
          </label>
          <select
            name="education"
            id="education"
            className="mt-1 block w-full rounded-sm border-gray-400 py-1 px-2"
          >
            <option value="SMA">SMA</option>
            <option value="S1">S1</option>
            <option value="S2">S2</option>
          </select>
        </div>

        <Input name="school" type="text" label="Sekolah/Universitas" />

        <button className="w-full rounded-sm bg-black py-2 text-white hover:bg-gray-800">
          Submit
        </button>
      </form>
    </div>
  );
}

export default RegisterPage;
