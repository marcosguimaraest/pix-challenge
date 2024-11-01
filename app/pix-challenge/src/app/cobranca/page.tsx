import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import FormCobranca from "../components/ui/form";
import Image from "next/image";

export default function Cobranca() {
    return (
        <div className="container mx-auto xl:p-14">
            <Image src={"data:image/pgn;base64, iVBORw0KGgoAAAANSUhEUgAAAcIAAAHCAQAAAABUY/ToAAADkElEQVR4Xu2VwW4bMRBDfdP//1E/a2/b8JHSOgGKFO2hU4CKvdJw+OiDxs7r/sP14/VV+d1V8rtV8rv1V+T1Yq2b07qWdlWLv2i32vHgLjmVtPzhicOmvILK60fcJeeSkgJx9OVfeLl2F3nFXXI4CSCYmu2keQQUUPJ/IW++yz7uswvmA23nlpxMct2wyCjUBCwRaUu3u+RYUjfO7f7WX9wlf/n3j8ln7btXDF2l4l6aA5SzSg4lZYo9wMXbN25EfR9l0qnkVNIdzYEI336KGJ8oGAn+ZpccSKoXCtNrJ8C7lXACw5ccSt6++3C2R5ItxkNxfOU/b8mRpE1nAuJ4M7pnIXLJuaQdl/oy7aC1U1zZxlhcX2ao5CgyopcD4JzimEflVXIsaezKAy8JdpxQeo+35GCStecBv+bg9emhTCGyrzMJJaeR0iWmsxnunsu3oCmQi+iLSSg5kbSdP7dObQu1RwXB6SWnkjjUt6IWk5CMPRX7ZE/JweS+ah3wPBF43dMHOBa55FxSfY+CPQ5Lyn1iIGwuOZhElwCobelNlGVwx0QrOZaM51y0qxhkz0l9F4oqOZbcC01t3pESJfaxrkzCW11yDKl7v8+dG+YlzRXD4ATrJWeT9uTsMg9d/E7g5A8qOZXct8tXVp18i/ckcAhjr7aSk0lRcgixIxI5Dt0Hn0sOJVViZn/O5jwCG3PJh5UcSj7ILUncLWvgZzY8Jmx3ybkkl2tdEfuM/QyCclI5uORQckNrX7M2tdmkkxgfVlUlh5LuLeqLScjRd58nA3FSrj1DJSeScgdfjIIrh2IQnzlQxnr7pS45kQTj1gWwFGcNQjUG+0sOJSXaZ9OHXV3cJ1MlbJSSo8nj+6D1dIySdHziDIOWHEpyMqDW6eEiM15/hq3+TSg5kNyYfGq4pwTZ1UhSIMSSc0k7/D3W42YGyHgSZfOHpCw5lEwbWZK7Jt6odL1f+aUuOZD0AFCm47NZvxwlSQ9tJYeSaoOs3RUS+w4iXQbn3CXnkjKHC/lSM8xOYR2t5GTSqrvB5XAK8I5C411yKClB2mV9OzlAEPZ20mwwCSVHkpLOEweFHodRvCnpJUeTmoF4ZNttD4gCPR9HQy45n4Slb0CMBEj53Se75HySHbcGIlduoxTH0Ps6CSUHkblxcXh1snxo2k/w+vQbX3IUSRdR/WXTLghiI8ddFSWHkn+4Sn63Sn63/gn5Ew4ojkpDXZX7AAAAAElFTkSuQmCC"
            } width={500} height={500} alt="Rei">

            </Image>
            <FormCobranca>
            </FormCobranca>
        </div>
    )
}