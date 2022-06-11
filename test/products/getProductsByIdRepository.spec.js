import { findById } from "../../src/products/repository";

const products = [
  {
    id: 1,
    name: "MacbookAir M2",
    desctiption: "New MacbookAir with M2 chip, incredibly thin and light",
    price: 1200,
    image: "m2air.jpg",
  },
  {
    id: 2,
    name: "MacbookPro M2",
    desctiption: "New MacbookPro with M2 chip, more powerful",
    price: 1500,
    image: "m2pro.jpg",
  },
];

jest.mock("../src/mocks/products.js", () => products);

describe("getProductsByIdRepository", () => {
  it("should return product by given id", () => {
    const product = findById(1);
    expect(product).toEqual(products[0]);
  });
});
