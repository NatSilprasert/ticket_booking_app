import { Button } from "@/components/Button";
import { VStack } from "@/components/VStack";
import { useAuth } from "@/context/AuthContext";

export default function SettingsScreen() {  
    const { logout } = useAuth();

    return (
        <VStack m={20}>
            <Button
                variant="outlined"
                onPress={() => logout()}
            >
                Logout
            </Button>
        </VStack>
    );
}