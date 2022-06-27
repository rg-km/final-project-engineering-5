import express from 'express';
import cors from 'cors';
import apiRoutes from './routes/index.js';

const PORT = process.env.PORT || 8080;

const app = express();
app.use(express.json());
app.use(cors());

app.use('/api', apiRoutes);

app.listen(PORT, () => {
  console.log(`Server running on: http://localhost:${PORT}`);
});
