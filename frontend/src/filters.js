const contains = (str, expr) => {
    return str.indexOf(expr) !== -1;
};

const truncate = str => str.length > 32 ? str.substring(0, 32) + '...'  : str; 

export default {
    contains,
    truncate
};