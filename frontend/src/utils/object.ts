export const sortObjectByValue = (obj: object): [string, any][] => {
    return Object.entries(obj).sort((a: any, b: any) => {
        return a[1] - b[1];
    })
}