import { AlertCircle } from "lucide-react";

interface CustomAlertProps {
    message: string;
    className?: string;
}

export default function CustomErrorAlert({ message, className }: CustomAlertProps) {
    return (
        <div className={className}>
            <span className="text-sm text-red-700 bg-red-100 py-3 px-4 rounded-sm w-full flex items-center">
                <AlertCircle className="text-red-700 mr-2" size="16" strokeWidth="3" />
                {message}
            </span>
        </div>
    )
}