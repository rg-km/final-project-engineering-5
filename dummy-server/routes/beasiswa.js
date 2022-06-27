import express from 'express';
import { beasiswa } from '../db/beasiswa.js';
import { paginate } from '../middleware/pagination.js';

const router = express.Router();

router.get('/', paginate('beasiswa', beasiswa), (_, res) => {
  res.json(res.paginatedResult);
});

export default router;
