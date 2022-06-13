import Input from './Input';

function MitraRegisterForm() {
  return (
    <form className="space-y-6 bg-gray-200 p-8">
      <Input name="name" type="text" label="Nama mitra" />

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

      <div>
        <label htmlFor="profile" className="block text-sm font-medium">
          Profil
        </label>
        <textarea
          name="profile"
          id="profile"
          rows={5}
          className="mt-1 block w-full rounded-sm border-gray-400 py-1 px-2"
        ></textarea>
      </div>

      <Input name="website" type="text" label="Situs" />

      <Input name="contactPerson" type="text" label="Contact person" />

      <button className="w-full rounded-sm bg-black py-2 text-white hover:bg-gray-800">
        Submit
      </button>
    </form>
  );
}

export default MitraRegisterForm;
