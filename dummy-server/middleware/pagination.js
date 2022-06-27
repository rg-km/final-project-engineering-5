export function paginate(route, model) {
  return (req, res, next) => {
    const page = Number(req.query.page) ?? 1;
    const limit = Number(req.query.limit) ?? 10;

    const startIndex = (page - 1) * limit;
    const endIndex = page * limit;

    const prevPage =
      startIndex > 0 ? `/api/${route}?page=${page - 1}&limit=${limit}` : '';
    const nextPage =
      endIndex < model.length
        ? `/api/${route}?page=${page + 1}&limit=${limit}`
        : '';

    res.paginatedResult = {
      data: model
        .slice(startIndex, endIndex)
        .filter((item) =>
          ['SISWA', 'MITRA'].includes(item.role)
            ? item.role.toLowerCase() === route
            : true
        ),
      prevPage,
      nextPage,
    };
    next();
  };
}
