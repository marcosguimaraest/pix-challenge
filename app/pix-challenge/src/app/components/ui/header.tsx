import Link from "next/link"
import Nav from "./nav";
import { Button } from "@/components/ui/button";

const Header = () => {

    return (
        <header className="py-8 xl:py12 text-purple-400">
            <div className="py-8 xl:py-12 container mx-auto flex justify-between items-center">
                <Link href="/">
                    <h1 className="text-4xl font-semibold">
                        PIXCHALLENGE<span className="text-slate-700">_</span>
                    </h1>
                </Link>
                <Nav />
                <div className="flex">
                    <Link href="https://github.com/marcosguimaraest">
                        <Button>
                            GITHUB
                        </Button>
                    </Link>
                </div>
            </div>
        </header>
    )
}

export default Header;