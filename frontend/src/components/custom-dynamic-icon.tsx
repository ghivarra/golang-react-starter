import { IconCollection } from "@/lib/icon-collection"

// dynamic icon interface
interface DynamicIconProps {
    name: keyof typeof IconCollection;
    color?: string;
    className?: string;
    size?: string | number;
    strokeWidth?: string | number;
}

// dynamic icon
export default function CustomDynamicIcon({name, size = 20, color = "black", ...props}: DynamicIconProps) {
    // find
    const IconComponent = IconCollection[name]

    // check if exist
    if (!IconComponent) {
        return null
    }

    return <IconComponent size={size} color={color} {...props} />
}