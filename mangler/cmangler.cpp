#include "mangler.hpp"
#include "cmangler.h"

Mangler NewMangler()
{
    auto mangler = new (codedweller::Mangler);
    return reinterpret_cast<Mangler>(mangler);
}

void DestroyMangler(Mangler mangler)
{
    delete reinterpret_cast<codedweller::Mangler *>(mangler);
}

unsigned char Encrypt(Mangler mangler, unsigned char c)
{
    return reinterpret_cast<codedweller::Mangler *>(mangler)->Encrypt(c);
}

unsigned char Decrypt(Mangler mangler, unsigned char c)
{
    return reinterpret_cast<codedweller::Mangler *>(mangler)->Decrypt(c);
}