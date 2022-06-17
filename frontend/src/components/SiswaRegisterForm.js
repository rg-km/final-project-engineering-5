import Input from './Input';

function SiswaRegisterForm({
  formValues,
  handleInputChange,
  showPassword,
  toggleShowPassword,
}) {
  return (
    <form className="space-y-6 bg-gray-200 p-8">
      <div className="grid grid-cols-2 gap-6">
        <Input
          name="firstName"
          value={formValues.firstName}
          type="text"
          label="Nama depan"
          onChange={handleInputChange}
        />
        <Input
          name="lastName"
          value={formValues.lastName}
          type="text"
          label="Nama belakang"
          onChange={handleInputChange}
        />
      </div>

      <Input
        name="email"
        value={formValues.email}
        type="email"
        label="Alamat email"
        onChange={handleInputChange}
      />

      <div className="grid grid-cols-2 gap-x-6 gap-y-2">
        <Input
          name="password"
          value={formValues.password}
          type={showPassword ? 'text' : 'password'}
          label="Password"
          onChange={handleInputChange}
        />
        <Input
          name="confirm"
          value={formValues.confirm}
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
        name="birthday"
        value={formValues.birthday}
        type="date"
        label="Tanggal lahir"
        onChange={handleInputChange}
      />

      <Input
        name="address"
        value={formValues.address}
        type="text"
        label="Alamat"
        onChange={handleInputChange}
      />

      <Input
        name="phoneNumber"
        value={formValues.phoneNumber}
        type="text"
        label="Nomor telepon"
        onChange={handleInputChange}
      />

      <Input
        name="accountNumber"
        value={formValues.accountNumber}
        type="text"
        label="Nomor rekening"
        onChange={handleInputChange}
      />

      <div>
        <label htmlFor="education" className="block text-sm font-medium">
          Tingkat pendidikan
        </label>
        <select
          name="education"
          value={formValues.education}
          defaultValue="default"
          id="education"
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
        name="school"
        value={formValues.school}
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
