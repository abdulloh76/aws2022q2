import { getAll, getById } from "./repository.js";

const getAllProducts = () => {
  return Promise.resolve(getAll());
};

const findProductById = (id) => {
  return Promise.resolve(getById(id));
};

export { getAllProducts, findProductById };
