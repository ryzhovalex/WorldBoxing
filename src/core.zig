pub const Codes = enum(i16){
    PersonCreated = -2,
    CreatePerson = -1,

    Ok = 0,

    Error = 1,
};
const Error = error{
    Default,
    UnrecognizedCommand,
    CommandParsing,
};