function LandingPageCard({ title, description, image }) {
  return (
    <div className="max-w-xs px-4 pt-6 text-center md:w-4/12 lg:pt-12">
      <div className="relative mb-8 flex h-64 w-full min-w-0 flex-col rounded-lg bg-white shadow-lg transition duration-300 ease-in-out hover:scale-110">
        <div className="flex-auto px-4 py-5">
          <div className="inline-flex w-auto items-center justify-center text-center">
            <img className="h-14" src={image} alt="" />
          </div>
          <div className="m-3">
            <h6 className="text-lg font-semibold">{title}</h6>
            <p className="text-blueGray-500 mt-2 mb-4 text-sm">{description}</p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default LandingPageCard;
