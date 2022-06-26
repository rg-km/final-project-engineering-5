import { format } from 'date-fns';
import { id } from 'date-fns/locale';
import { useEffect, useState } from 'react';
import shallow from 'zustand/shallow';
import { applyBeasiswa } from '../lib/beasiswa';
import { getSiswa } from '../lib/siswa';
import useAuthStore from '../store/auth';

function BeasiswaDetailSiswa({ beasiswa, mitra }) {
  const { user, isAuthenticated } = useAuthStore(
    (state) => ({ user: state.user, isAuthenticated: state.isAuthenticated }),
    shallow
  );
  const [terdaftar, setTerdaftar] = useState(false);

  useEffect(() => {
    if (!isAuthenticated || user.role !== 'SISWA') {
      return;
    }
    if (beasiswa) {
      getSiswa(user.token).then((res) => {
        setTerdaftar(!!res.data.find((item) => item.id === beasiswa.id));
      });
    }
  }, [isAuthenticated, beasiswa, user.token]);

  return (
    <div className="rounded-lg border border-gray-300 p-4 shadow-md">
      {!beasiswa ? (
        <p className="text-center">
          Pilih beasiswa untuk melihat info lebih detail
        </p>
      ) : (
        <>
          <h2 className="text-xl font-semibold">
            {beasiswa.judulBeasiswa || beasiswa.namaBeasiswa}
          </h2>
          <div className="flex gap-4">
            <p>{beasiswa.namaMitra || 'Mitra'}</p>
            {mitra && (
              <p className="font-semibold">
                {mitra.pic} ({mitra.nomorPic})
              </p>
            )}
          </div>
          {beasiswa.tanggalPembukaan && beasiswa.tanggalPenutupan && (
            <div className="mt-2 grid w-fit grid-cols-[auto_auto] gap-x-4">
              <span className="font-semibold">Buka</span>
              <time dateTime={beasiswa.tanggalPembukaan}>
                {format(new Date(beasiswa.tanggalPembukaan), 'dd MMMM yyyy', {
                  locale: id,
                })}
              </time>
              <span className="font-semibold">Tutup</span>
              <time dateTime={beasiswa.tanggalPenutupan}>
                {format(new Date(beasiswa.tanggalPenutupan), 'dd MMMM yyyy', {
                  locale: id,
                })}
              </time>
            </div>
          )}
          {(beasiswa.status || beasiswa.statusPendaftaran) && (
            <div
              className={`mt-8 inline-block rounded border px-2 py-1 text-sm font-semibold ${
                (beasiswa.status || beasiswa.statusPendaftaran) === 'Diterima'
                  ? 'bg-green-100 text-green-900'
                  : (beasiswa.status || beasiswa.statusPendaftaran) ===
                    'Ditolak'
                  ? 'bg-red-100 text-red-900'
                  : 'bg-gray-100 text-amber-500'
              }`}
            >
              {beasiswa.status || beasiswa.statusPendaftaran}
            </div>
          )}
          <p className="mt-4">{beasiswa.deskripsi}</p>
          {isAuthenticated &&
            (terdaftar ? (
              <p className="mt-4 font-medium">Terdaftar</p>
            ) : (
              user.role === 'SISWA' && (
                <button
                  className="mt-4 rounded border border-transparent bg-black px-4 py-1 text-white hover:bg-gray-800"
                  onClick={() => {
                    try {
                      applyBeasiswa(user.token, user.idSiswa, beasiswa.id);
                    } catch (error) {
                      console.log(error.message);
                    }
                  }}
                >
                  Daftar
                </button>
              )
            ))}
        </>
      )}
    </div>
  );
}

export default BeasiswaDetailSiswa;
