function TableSiswaComponent({ dataSiswa }) {
  return (
    <div className="flex flex-col">
      <div className="overflow-x-auto">
        <div className="inline-block min-w-full py-2">
          <div className="overflow-hidden rounded-md">
            <table className="min-w-full">
              <thead className="border-b bg-slate-200">
                <tr>
                  <th className="px-6 py-4 text-left text-sm font-medium text-gray-900">
                    No
                  </th>
                  <th className="px-6 py-4 text-left text-sm font-medium text-gray-900">
                    Nama
                  </th>
                  <th className="px-6 py-4 text-left text-sm font-medium text-gray-900">
                    Instansi
                  </th>
                  <th className="px-6 py-4 text-left text-sm font-medium text-gray-900">
                    Alamat
                  </th>
                  <th className="px-6 py-4 text-left text-sm font-medium text-gray-900">
                    Email/No Hp
                  </th>
                  <th className="px-6 py-4 text-left text-sm font-medium text-gray-900">
                    Status
                  </th>
                </tr>
              </thead>
              <tbody>
                {dataSiswa.map((data, index) => (
                  <tr
                    key={index}
                    className="border-b bg-white transition duration-300 ease-in-out hover:bg-gray-100"
                    onClick={() => console.log('Clicked ' + data.id)}
                  >
                    <td className="whitespace-nowrap px-6 py-4 text-sm font-medium text-gray-900">
                      {data.id}
                    </td>
                    <td className="whitespace-nowrap px-6 py-4 text-sm font-normal text-gray-900">
                      {data.nama}
                    </td>
                    <td className="whitespace-nowrap px-6 py-4 text-sm font-normal text-gray-900">
                      {data.namaInstansi}
                    </td>
                    <td className="whitespace-nowrap px-6 py-4 text-sm font-normal text-gray-900">
                      {data.alamat}
                    </td>
                    <td className="whitespace-nowrap px-6 py-4 text-sm font-normal text-gray-900">
                      {data.email}
                    </td>
                    <td className="whitespace-nowrap px-6 py-4 text-sm font-normal text-gray-900">
                      {data.status}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
}

export default TableSiswaComponent;
