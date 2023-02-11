import AppRouter from './AppRouter'
import { AuthProvider } from './contextAPI/AuthProvider'
import "./scss/app.scss"

function App() {
  return (
    <div className="App">
      <AuthProvider>
        <AppRouter />
      </AuthProvider>
    </div>
  );
}

export default App;
