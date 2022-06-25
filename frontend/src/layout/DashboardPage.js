import { useEffect, useState } from 'react';
import ClipLoader from 'react-spinners/ClipLoader';
import BeasiswaDetail from '../components/BeasiswaDetail';
import BeasiswaList from '../components/BeasiswaList';
import TableSiswaComponent from '../components/TableSiswaComponent';
import useAuthStore from '../store/auth';

const detailSiswa = {
  siswa: {
    id: 1,
    nama: 'John Doe',
    tanggalLahir: '2000-01-02',
    email: 'johndoe@email.com',
    tingkatPendidikan: 'S1',
    namaInstansi: 'ITB',
    nomorRekening: '123456',
    namaBank: 'BCA',
  },
  dataBeasiswa: [
    {
      id: 1,
      nama: 'Beasiswa Bakti BCA',
      IdMitra: 1,
      namaMitra: 'BCA',
      statusPendaftaran: 'Diterima',
    },
    {
      id: 2,
      nama: 'Beasiswa ABB Jurgen Dormann Foundation',
      IdMitra: 2,
      namaMitra: 'ABB Jurgen Dormann Foundation',
      statusPendaftaran: 'Ditolak',
    },
  ],
};

const detailMitra = {
  nama: 'BCA',
  email: 'halobca@bca.co.id',
  about: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit',
  nomorPic: '081212345678',
  situs: 'http://www.bca.co.id',
  pic: 'Jane Doe',
};

const beasiswaMitraList = [
  {
    nama: 'Beasiswa Bakti BCA',
    IdMitra: '1',
    namaMitra: 'BCA',
    benefits: '400000',
    deskripsi: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
    tanggalPembukaan: '2022-06-17',
    tanggalPenutupan: '2022-07-17',
    jumlahPendaftar: 57,
  },
  {
    nama: 'Beasiswa Bakti BCA 2',
    IdMitra: '1',
    namaMitra: 'BCA',
    benefits: '500000',
    deskripsi: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
    tanggalPembukaan: '2022-08-24',
    tanggalPenutupan: '2022-09-24',
    jumlahPendaftar: 21,
  },
];

function DashboardPage() {
  const [detailUser, setDetailUser] = useState(null);
  const [activeBeasiswa, setActiveBeasiswa] = useState(null);
  const user = useAuthStore((state) => state.user);

  useEffect(() => {
    setTimeout(() => {
      if (user.role === 'SISWA') {
        setDetailUser(detailSiswa);
      } else if (user.role === 'MITRA') {
        setDetailUser({
          mitra: detailMitra,
          dataBeasiswa: beasiswaMitraList,
        });
      }
    }, 500);
  }, []);

  return (
    <div>
      {!detailUser ? (
        <div className="mt-16 text-center">
          <ClipLoader />
        </div>
      ) : (
        <div className="mx-auto grid max-w-screen-lg grid-cols-[1fr_2fr] items-start gap-4 py-16 px-4">
          <BeasiswaList
            beasiswaList={detailUser.dataBeasiswa}
            activeBeasiswa={activeBeasiswa}
            setActiveBeasiswa={setActiveBeasiswa}
          />
          {detailUser.dataBeasiswa && (
            <BeasiswaDetail
              beasiswa={activeBeasiswa}
              mitra={detailMitra}
              detailSiswa={detailSiswa}
            />
            /* <TableSiswaComponent detailSiswa={detailUser} />
            </BeasiswaDetail> */
          )}
        </div>
      )}
    </div>
  );
}

export default DashboardPage;