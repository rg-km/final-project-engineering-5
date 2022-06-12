function LoginPage() {
  return (
    <section className="h-screen">
      <div className="container h-full px-6 py-12">
        <div className="g-6 flex h-full flex-wrap items-center justify-center text-gray-800">
          {/* Image Jika Butuh*/}
          <div className="mb-12 md:mb-0 md:w-8/12 lg:w-6/12">
            <img className="w-full" src="" alt="" />
          </div>

          <div className="block rounded-xl bg-white py-12 px-10 shadow-lg md:w-8/12 lg:ml-20 lg:w-5/12">
            <form>
              <h1 className="mb-8 text-3xl font-semibold">Selamat Datang</h1>

              {/* Email Input */}
              <div className="mb-6">
                <input
                  className="form-control m-0 block w-full rounded border border-solid border-gray-300 bg-white bg-clip-padding px-4 py-2 text-xl font-normal text-gray-700 transition ease-in-out focus:bg-white focus:text-gray-700 focus:outline-none focus:ring focus:ring-blue-600"
                  type="text"
                  placeholder="Email address"
                />
              </div>

              {/* Password Input */}
              <div className="mb-6">
                <input
                  className="form-control m-0 block w-full rounded border border-solid border-gray-300 bg-clip-padding px-4 py-2 text-xl font-normal text-gray-700 transition ease-in-out focus:bg-white focus:text-gray-700 focus:outline-none focus:ring focus:ring-blue-600"
                  type="password"
                  placeholder="Password"
                />
              </div>

              {/* CheckBox */}
              <div className="mb-6 flex items-center justify-between">
                <div className="form-check">
                  <input
                    type="checkbox"
                    className="form-checkbox float-left mt-1 mr-2 h-4 w-4 cursor-pointer appearance-none rounded-sm border border-gray-300 bg-white bg-contain bg-center bg-no-repeat align-top transition duration-200 checked:border-blue-600 checked:bg-blue-600 focus:outline-none focus:ring-0"
                  />
                  <label className="form-check-label inline-block font-semibold text-gray-800">
                    Remember me
                  </label>
                </div>
                <a
                  href="#!"
                  className="font-semibold text-blue-600 transition duration-200 ease-in-out hover:text-blue-700 focus:text-blue-700 active:text-blue-800"
                >
                  Forgot password?
                </a>
              </div>

              {/* Login Button */}
              <button
                type="submit"
                className="inline-block w-full rounded bg-blue-600 px-7 py-3 text-sm font-medium uppercase leading-snug text-white shadow-md transition duration-150 ease-in-out hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg"
                data-mdb-ripple="true"
                data-mdb-ripple-color="light"
              >
                Login
              </button>

              <div className="my-4 flex items-center rounded before:mt-2 before:flex-1 before:border before:border-gray-300"></div>

              <div className="text-center">
                <p className="mt-2 mb-0 pt-1 text-sm font-semibold">
                  Belum punya akun?{' '}
                  <a
                    href="#!"
                    className=" text-red-600 transition duration-200 ease-in-out hover:text-red-700 focus:text-red-700"
                  >
                    Daftar
                  </a>
                </p>
              </div>
            </form>
          </div>
        </div>
      </div>
    </section>
  );
}

export default LoginPage;
