import '../styles/ProductList.css';
import ProductCard from './ProductCard';

// function ProductList({ products, onEdit, onDelete }) {
//   return (
//     <div className="product-list">
//       {products.map((product) => (
//         <ProductCard
//           key={product._id}
//           product={product}
//           onEdit={onEdit}
//           onDelete={onDelete}
//         />
//       ))}
//     </div>
//   );
// }


// ProductList.js
function ProductList({ products, onEdit, onDelete }) {
    console.log(products)
    return (
        <div className="product-list">
            {products.map((product) => (
                <ProductCard
                    key={product.id}
                    product={product}
                    onEdit={() => onEdit(product)} // Ensure onEdit is a function that handles editing
                    onDelete={onDelete}
                />
            ))}
        </div>
    );
}

export default ProductList;