import DetailTableComponent from '../components/DetailTableComponent';

function DetailSiswaPage() {
  return (
    <div className="my-4 mx-16 overflow-hidden rounded-lg shadow">
      <div className="px-4 py-5">
        <h3 className="text-lg font-medium leading-6 text-gray-900">
          Detail Pendaftar
        </h3>
        <p className="mt-1 max-w-2xl text-sm text-gray-500">
          Berikut adalah detail pendaftar.
        </p>
      </div>
      <div className="border-t border-gray-200">
        <dl>
          <DetailTableComponent />
          <div className="flex justify-end space-x-4 bg-gray-50 px-4 py-5">
            <button
              type="button"
              class="inline-block rounded bg-green-500 px-6 py-2.5 text-xs font-medium uppercase leading-tight text-white shadow-md transition duration-150 ease-in-out hover:bg-green-600 hover:shadow-lg focus:bg-green-600 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-green-700 active:shadow-lg"
            >
              Terima
            </button>
            <button
              type="button"
              class="inline-block rounded bg-red-600 px-6 py-2.5 text-xs font-medium uppercase leading-tight text-white shadow-md transition duration-150 ease-in-out hover:bg-red-700 hover:shadow-lg focus:bg-red-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-red-800 active:shadow-lg"
            >
              Tolak
            </button>
          </div>
        </dl>
      </div>
    </div>
  );
}

export default DetailSiswaPage;
