import { useEffect, useState } from 'react';
import ClipLoader from 'react-spinners/ClipLoader';
import BeasiswaDetail from '../components/BeasiswaDetail';
import BeasiswaList from '../components/BeasiswaList';
import { getBeasiswaList } from '../lib/beasiswa';

function BeasiswaListPage() {
  const [beasiswaList, setBeasiswaList] = useState(null);
  const [activeBeasiswa, setActiveBeasiswa] = useState(null);

  useEffect(() => {
    getBeasiswaList().then((d) => {
      setBeasiswaList(d.data);
    });
  }, []);

  return (
    <div>
      {!beasiswaList ? (
        <div className="mt-16 text-center">
          <ClipLoader />
        </div>
      ) : (
        <div className="mx-auto grid max-w-screen-lg grid-cols-[1fr_2fr] items-start gap-4 py-16 px-4">
          <BeasiswaList
            beasiswaList={beasiswaList}
            activeBeasiswa={activeBeasiswa}
            setActiveBeasiswa={setActiveBeasiswa}
          />
          {beasiswaList && (
            <BeasiswaDetail
              beasiswa={activeBeasiswa}
              mitra={{ pic: 'John Doe', nomorPic: '081212345678' }}
            />
          )}
        </div>
      )}
    </div>
  );
}

export default BeasiswaListPage;
