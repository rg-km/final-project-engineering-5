import { useState } from 'react';
import Input from '../components/Input';

function LoginPage() {
  const [formValues, setFormValues] = useState({});

  const handleInputChange = (e) => {
    setFormValues({
      ...formValues,
      [e.target.name]: e.target.value,
    });
  };

  return (
    <div className="mx-auto max-w-[448px] py-10">
      <form className="space-y-6 bg-gray-200 p-8">
        <Input
          name="email"
          value={formValues.email}
          type="email"
          label="Alamat email"
          onChange={handleInputChange}
        />
        <Input
          name="password"
          value={formValues.password}
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
