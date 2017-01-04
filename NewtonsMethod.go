package main

import(
    "math" 
)

const EPS = 0.000001;
const MIN = 0.000001;

func IsEqual(f1, f2 float64) bool {
    return math.Abs(f1 -f2) < MIN
}

func IsZeroFloat(f1 float64) bool {
    return f1 == float64(0);
}

func Sqrt2(fValue float64) float64 {
    if IsZeroFloat(fValue) {
         return fValue;
    }

    Result := fValue;
    LastValue := float64(0);

    for {
        LastValue = Result;
        Result = Result / 2.0 +  fValue / Result / 2.0;
        if IsEqual(LastValue, Result) {
            break;
        }
    }
    return Result;
}

func FuncX(fValue float64) float64 {
    return fValue * fValue - 2; // 对应求方程x * x = 2 的x的解。
}

func DerivedFuncX(fValue float64) float64 {
    return 2 * fValue;
}

func Sqrt2Other(fValue float64) float64 {
    if IsZeroFloat(fValue) {
        return fValue;
    }

    x0 := fValue;
    x1 := float64(0.0);

    for {
        x1 = x0 - FuncX(x0)/DerivedFuncX(x0);
        if IsEqual(x1, x0) {
            break;
        } else {
            x0 = x1;
        }
    }
    return x1;
}