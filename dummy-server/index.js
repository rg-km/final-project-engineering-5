import express from 'express';
import apiRoutes from './routes/index.js';

const PORT = process.env.PORT || 5000;

const app = express();
app.use(express.json());

app.use('/api', apiRoutes);

app.listen(PORT, () => {
  console.log(`Server running on: http://localhost:${PORT}`);
});
