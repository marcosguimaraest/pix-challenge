import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

export default function FormCobranca() {
    return (<form action="">
        <Label htmlFor="customerName">Nome do cliente</Label>
        <Input type="text" name="customerName" id="customerName" placeholder="_nome do cliente" />

        <Label htmlFor="customerCpf">CPF do cliente</Label>
        <Input type="text" name="customerCpf" id="Cpf" placeholder="_cpf do cliente" />

        <Label htmlFor="value">Valor da cobrança</Label>
        <Input type="number" name="value" id="value" placeholder="_valor da cobranca" />

        <Input className="hidden" type="text" readOnly name="operationType" id="operationType" value="PIX" />

        <Label htmlFor="date">Data de expiração</Label>
        <Input type="date" name="date" id="date" placeholder="_date" />

        <Button>
            Submit
        </Button>
    </form>)
}