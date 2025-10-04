import { DimensionValue } from "react-native";

export interface ShortcutProps {
    m?: number | "auto";
    ml?: number | "auto";
    mr?: number | "auto";
    mt?: number | "auto";
    mb?: number | "auto";
    mx?: number | "auto";
    my?: number | "auto";

    p?: number;
    pl?: number;
    pr?: number;
    pt?: number;
    pb?: number;
    px?: number;
    py?: number;

    w?: DimensionValue;
    h?: DimensionValue;
}

export const defaultShortcuts = (props: ShortcutProps) => ({
    // Margin
    margin: props.m,
    marginLeft: props.ml ?? props.mx,
    marginRight: props.mr ?? props.mx,
    marginTop: props.mt ?? props.my,
    marginBottom: props.mb ?? props.my,

    // Padding
    padding: props.p,
    paddingLeft: props.pl ?? props.px,
    paddingRight: props.pr ?? props.px,
    paddingTop: props.pt ?? props.py,
    paddingBottom: props.pb ?? props.py,

    // Width & Height
    width: props.w,
    height: props.h,
});